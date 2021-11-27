require('dotenv').config();
import { spawn } from "child_process";
import { mkdirSync, existsSync, writeFileSync } from "fs";
import { downloadInputForYearAndDay, getPuzzleDescription, getTestContents } from './utils/aoc-actions';
import { cp } from 'shelljs';

const action = process.argv[2]
const year = process.argv[3]
const day = process.argv[4]


const createFromTemplate = async () => {
  let path = `./challenges/${year}/${day}`;
  if (!existsSync(path)) {
    console.log(`Creating challenge to ${path} from template...`);
    mkdirSync(`challenges/${year}/${day}`, { recursive: true});
    //Copy template
    cp('-rf', "template/*", path);
  }

  if (!existsSync(`${path}/input.txt`)) {
    console.log(`Downloading input...`);
    let input = await downloadInputForYearAndDay(day, year);
    // @ts-ignore
    writeFileSync(`${path}/input.txt`, input)
  }
  let readme = await getPuzzleDescription(year, day);
  // @ts-ignore
  writeFileSync(`${path}/README.html`, readme)

  let tests = getTestContents(year, day);
  writeFileSync(`${path}/index.test.ts`, tests)
}

if (action === 'create') {
  createFromTemplate();
}

if (action === 'run') {
  const path = `challenges/${year}/${day}/index.ts`;
  if (existsSync(path)) {
    spawn("nodemon", ["-x", "ts-node", `challenges/${year}/${day}/index.ts ${year} ${day}`], {
      stdio: "inherit",
      shell: true,
    })
  }
}
