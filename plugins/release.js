const {Plugin} = require('release-it');
const {Octokit} = require('@octokit/rest');
const pkg = require('../package.json');
const fs = require('fs');

class WandbPlugin extends Plugin {
  async init() {
    this.octokit = new Octokit({
      baseUrl: 'https://api.github.com',
      auth: process.env['GITHUB_TOKEN'],
      userAgent: `local/${pkg.version}`,
      log: null,
      request: {
        timeout: 10000,
      },
    });
    this.registerPrompts({
      release_notes: {
        type: 'editor',
        name: 'release_notes',
        default: this.defaultNotes.bind(this),
        message: () => "Let's edit the release notes.",
      },
    });
  }

  defaultNotes() {
    return this.getContext("notes");
  }

  bump(version) {
    this.setContext({version});
  }

  async getChangelog(latestVersion) {
    this.setContext({latestVersion});
    if (this.options.legacy) {
      const versionParts = latestVersion.split(".").map((v) => parseInt(v, 10))
      versionParts[2] += 1
      const version = versionParts.join(".")
      const res = await this.octokit.repos.getReleaseByTag({
        owner: 'wandb',
        repo: 'core',
        tag: `local/v${version}`,
      });
      this.setContext({date: new Date(res.data.published_at), notes: res.data.body});
      return res.data.body;
    } else {
      this.setContext({date: new Date()});
      let res = await this.octokit.repos.getReleaseByTag({
        owner: 'wandb',
        repo: 'core',
        tag: `local/v${latestVersion}`,
      });
      const lastReleaseSHA = res.data.target_commitish;
      const publishedAt = res.data.published_at;
      console.log('Grabbing all commits since', publishedAt, lastReleaseSHA);
      res = await this.octokit.repos.listCommits({
        owner: 'wandb',
        repo: 'core',
        per_page: 100,
        since: publishedAt,
      });
      //new Date(commit.author.date) > new Date(publishedAt)
      if(res.data.length > 100) {
        console.warn("There have been more than 100 commits since the last release!")
      }
      const notes = res.data
        .map((commit) => `* ${commit.commit.message.split('\n')[0]}`)
        .join('\n');
      this.setContext({notes})
      return notes;
    }
  }

  saveChangelogToFile(filePath, renderedTemplate) {
    const fileDescriptor = fs.openSync(filePath, 'a+');

    const oldData = fs.readFileSync(filePath);
    const newData = new Buffer.from(renderedTemplate.split("\r\n").join("\n"));

    fs.writeSync(fileDescriptor, newData, 0, newData.length, 0);
    fs.writeSync(fileDescriptor, oldData, 0, oldData.length, newData.length);

    fs.closeSync(fileDescriptor);
  }

  saveReleaseNotesToFile(filePath, notes) {
    const fileDescriptor = fs.openSync(filePath, 'a+');
    const newData = new Buffer.from(notes.split("\r\n").join("\n"));
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
      task: (notes) => {
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
}

module.exports = WandbPlugin;