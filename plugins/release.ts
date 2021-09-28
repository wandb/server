import {Plugin} from 'release-it';
import {Octokit} from '@octokit/rest';
import pkg from '../package.json';
import fs from 'fs';
import * as autoReleaseNotes from 'auto-release-notes/lib/index';
import {compact} from 'lodash';

const getReleaseNotesBuffer = (
  releaseNotes: string[],
  commitsToInspect: autoReleaseNotes.Commit[]
) => {
  const formattedReleaseNotes = releaseNotes.join('\n* ');
  const formattedCommitsToInspect = commitsToInspect
    .map(
      (commit) => `* ${commit.commit.message.split('\n')[0]} ${commit.html_url}`
    )
    .join('\n');

  return `* ${formattedReleaseNotes}


${
  commitsToInspect.length > 0
    ? `
## Ambiguous Commits

Release notes could not be inferred for the following commits -- please check them manually, and then remove this section.


${formattedCommitsToInspect}`
    : ''
}
`;
};

class WandbPlugin extends Plugin {
  async init() {
    this.octokit = new Octokit({
      baseUrl: 'https://api.github.com',
      auth: process.env['GITHUB_TOKEN'],
      userAgent: `local/${pkg.version}`,
      log: null as any,
      request: {
        timeout: 10000,
      },
    });

    this.setContext({date: new Date()});

    this.registerPrompts({
      release_notes: {
        type: 'editor',
        name: 'release_notes',
        default: this.defaultNotes.bind(this),
        message: () => "Let's edit the release notes.",
      },
      release_to_core: {
        type: 'confirm',
        name: 'release_to_core',
        message: () =>
          "Let's create a release in core to trigger the push to dockerhub",
      },
    });
  }

  defaultNotes() {
    return this.getContext('notes');
  }

  bump(version: string) {
    this.setContext({version});
  }

  async getChangelog(latestVersion: string) {
    this.setContext({latestVersion});
    if (this.options.legacy) {
      const versionParts = latestVersion.split('.').map((v) => parseInt(v, 10));
      versionParts[2] += 1;
      const version = versionParts.join('.');
      const res = await this.octokit.repos.getReleaseByTag({
        owner: 'wandb',
        repo: 'core',
        tag: `local/v${version}`,
      });
      this.setContext({
        date: new Date(res.data.published_at),
        notes: res.data.body,
      });
      return res.data.body;
    } else {
      const {commitsSinceLastRelease} =
        await autoReleaseNotes.getLastReleaseInfo(
          this.octokit,
          'wandb',
          'core',
          `local/v${latestVersion}`
        );

      const lastFiveCommitChoices = commitsSinceLastRelease
        .slice(0, 5)
        .map((commit) => ({
          name: `${commit.sha.slice(0, 8)} - ${commit.commit.message
            .split('\n')[0]
            .slice(0, 100)}`,
          value: commit.sha, // this will be returned as the choice value
          short: commit.sha.slice(0, 8),
        }));

      this.registerPrompts({
        select_target_commit: {
          type: 'list',
          name: 'select_target_commit',
          default: 0,
          message: () => "Let's select which commit to release",
          choices: () => lastFiveCommitChoices,
          pageSize: 5,
        },
      });

      await this.step({
        enabled: true,
        task: (targetSHA: string) => {
          this.setContext({targetSHA});
        },
        label: 'Selecting target commit',
        prompt: 'select_target_commit',
      });
      const {targetSHA} = this.getContext();

      const targetSHAIndex = commitsSinceLastRelease.findIndex(
        (commit) => commit.sha === targetSHA
      );

      const sleep = (time: number) =>
        new Promise((resolve) => {
          setTimeout(() => {
            resolve('resolved');
          }, time);
        });

      const commitsWithReleaseNotes: Array<
        [autoReleaseNotes.Commit, string[] | null]
      > = [];

      for (const commit of commitsSinceLastRelease) {
        await sleep(100);
        const notes = await autoReleaseNotes.getReleaseNotesForCommit(
          this.octokit,
          'wandb',
          'core',
          commit
        );
        commitsWithReleaseNotes.push([commit, notes]);
      }

      // const commitsWithReleaseNotes = await Promise.all(
      //   commitsSinceLastRelease.map(
      //     async (commit) =>
      //       [
      //         commit,
      //         autoReleaseNotes.getReleaseNotesForCommit(
      //           this.octokit,
      //           'wandb',
      //           'core',
      //           commit
      //         ),
      //       ] as const
      //   )
      // );

      const allReleaseNotes = compact(
        commitsWithReleaseNotes.map(([_, releaseNotes]) => releaseNotes)
      ).flat();

      const commitsToInspect = commitsWithReleaseNotes
        .filter(([_, releaseNotes]) => releaseNotes == null)
        .map(([commit, _]) => commit);

      const notes = getReleaseNotesBuffer(allReleaseNotes, commitsToInspect);

      this.setContext({notes});

      return notes;
    }
  }

  saveChangelogToFile(filePath: string, renderedTemplate: string) {
    const fileDescriptor = fs.openSync(filePath, 'a+');

    const oldData = fs.readFileSync(filePath);
    const newData = Buffer.from(renderedTemplate.split('\r\n').join('\n'));

    fs.writeSync(fileDescriptor, newData, 0, newData.length, 0);
    fs.writeSync(fileDescriptor, oldData, 0, oldData.length, newData.length);

    fs.closeSync(fileDescriptor);
  }

  saveReleaseNotesToFile(filePath: string, notes: string) {
    const fileDescriptor = fs.openSync(filePath, 'a+');
    const newData = Buffer.from(notes.split('\r\n').join('\n'));
    fs.writeSync(fileDescriptor, newData, 0, newData.length, 0);
    fs.closeSync(fileDescriptor);
  }

  getFormattedDate() {
    const date = this.getContext('date');
    return date.toLocaleDateString('en-us', {
      month: 'long',
      day: 'numeric',
      year: 'numeric',
    });
  }

  async beforeRelease() {
    await this.step({
      enabled: true,
      task: (notes: string) => {
        this.setContext({notes});
        this.saveReleaseNotesToFile('staging/RELEASE.md', notes);
        this.saveChangelogToFile(
          'CHANGELOG.md',
          `## wandb/local:${this.getContext(
            'version'
          )} (${this.getFormattedDate()})\n\n${notes}\n\n`
        );
      },
      label: 'Creating notes',
      prompt: 'release_notes',
    });
  }

  async release() {
    await this.step({
      enabled: !this.options.legacy,
      task: async () => {
        await this.octokit.repos.createRelease({
          owner: 'wandb',
          repo: 'core',
          tag_name: `local/v${this.getContext('version')}`,
          name: `Local v${this.getContext('version')}`,
          body: this.getContext('notes'),
          target_commitish: this.getContext('targetSHA'),
        });
      },
      label: 'Creating release in core',
      prompt: 'release_to_core',
    });
  }
}

module.exports = WandbPlugin;
