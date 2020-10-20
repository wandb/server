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

  async defaultNotes() {
    if (this.options.legacy) {
      // TODO: version
      const res = await this.octokit.repos.getReleaseByTag({
        owner: 'wandb',
        repo: 'core',
        tag: 'local/v0.9.4',
      });
      this.setContext({date: new Date(res.data.published_at)});
      return res.data.body;
    } else {
      this.setContext({date: new Date()});
      // TODO: latestVersion
      let res = await this.octokit.repos.getReleaseByTag({
        owner: 'wandb',
        repo: 'core',
        tag: 'local/v0.9.30',
      });
      const lastReleaseSHA = res.data.target_commitish;
      const publishedAt = res.data.published_at;
      console.log('Grabbing all commits from since', publishedAt, lastReleaseSHA);
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
      return notes;
    }
  }

  bump(version) {
    this.setContext({version});
  }

  getChangelog(latestVersion) {
    this.setContext({latestVersion});
  }

  saveChangelogToFile(filePath, renderedTemplate) {
    const fileDescriptor = fs.openSync(filePath, 'a+');

    const oldData = fs.readFileSync(filePath);
    const newData = new Buffer.from(renderedTemplate);

    fs.writeSync(fileDescriptor, newData, 0, newData.length, 0);
    fs.writeSync(fileDescriptor, oldData, 0, oldData.length, newData.length);

    fs.closeSync(fileDescriptor);
  }

  saveReleaseNotesToFile(filePath, notes) {
    const fileDescriptor = fs.openSync(filePath, 'a+');
    const newData = new Buffer.from(notes);
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
        console.log('CTX', this.getContext());
        this.saveReleaseNotesToFile('staging/RELEASE.md', notes);
        this.saveChangelogToFile(
          'CHANGELOG.md',
          `# wandb/local:${this.getContext(
            'version'
          )} - ${this.getFormattedDate()}\n\n${notes}\n\n`
        );
      },
      label: 'Creating notes',
      prompt: 'release_notes',
    });
  }
}

module.exports = WandbPlugin;
