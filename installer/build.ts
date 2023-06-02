import * as fs from 'fs'
// import * as readline from 'readline'
import { spawnSync } from 'child_process'

const INSTALLER_SCRIPT = `${__dirname}/install.sh`
const INSTALLER_SCRIPT_VARIABLE_DIR = 'DIR'

const generateAirgapScript = Boolean(process.env.AIRGAP)

function readBashVariables(filePath: string) {
  const content = fs.readFileSync(filePath, { encoding: 'utf-8' })
  const lines = content.split('\n')
  const result: Record<string, string | undefined> = {}
  for (const line of lines) {
    const l = line.startsWith('export ') ? line.slice(7) : line
    const match = l.match(/^\s*([\w-]+)\s*=\s*(.*)\s*/)
    if (match) {
      const variable = match[1]
      const value = match[2].replace(/^"|^'|"$/g, '')
      const output = spawnSync('bash', ['-c', `echo ${value}`])

      result[variable] = output.stdout.toString().trim()
    }
  }
  return result
}

function getImportLines(filePath: string) {
  const content = fs.readFileSync(filePath, { encoding: 'utf-8' })
  const lines = content.split('\n')
  const imports: string[] = []
  let withinTags = false
  for (const line of lines) {
    if (line.startsWith('# </ImportInline>')) withinTags = false
    if (withinTags && !line.startsWith('#')) imports.push(line)
    if (line.includes('# <ImportInline>')) withinTags = true
  }
  return imports
}

function main() {
  const scriptVariables = readBashVariables(INSTALLER_SCRIPT)
  const dir = scriptVariables[INSTALLER_SCRIPT_VARIABLE_DIR]

  if (dir == null) {
    console.error(
      `Could not find '${INSTALLER_SCRIPT_VARIABLE_DIR}' variable in bash script`,
    )
    return
  }

  let content = fs
    .readFileSync(INSTALLER_SCRIPT, { encoding: 'utf-8' })
    .split('\n')

  if (generateAirgapScript)
    content = content.map((d) =>
      d.startsWith('export AIRGAP=') ? 'export AIRGAP=1' : d,
    )

  for (const importLine of getImportLines(INSTALLER_SCRIPT)) {
    const importScriptPath = importLine.replace(
      `. $${INSTALLER_SCRIPT_VARIABLE_DIR}`,
      dir,
    )
    console.log(`Importing: ${importScriptPath} (${importLine})`)
    const importScript = fs.readFileSync(importScriptPath, {
      encoding: 'utf-8',
    })

    const i = content.findIndex((l) => l === importLine)
    if (i === -1) continue
    content[i] = `${importScript}\n`
  }

  fs.writeFileSync('./install.sh', content.join('\n'))
}

main()
