import * as fs from 'fs'

const SCRIPT_PATH = `${__dirname}/install.sh`

const getImportLines = () => {
  const input = fs.createReadStream(SCRIPT_PATH)
  const lines = readline.createInterface({
    input,
    crlfDelay: Infinity,
  })
  for await (const line of lines) {
    if (line.contains('</Inline>')) {
      withinTags = false
      if (withinTags) {
        inlineValues.push(line)
      }

      if (line.contains('<Inline>')) {
        withinTags = true
      }
    }
  }
}
