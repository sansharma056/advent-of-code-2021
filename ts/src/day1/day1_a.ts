import * as fs from "fs";
import readline from "readline";

async function countNumOfIncDepths(): Promise<number> {
	let count = 0;

	const rl = readline.createInterface({
		input: fs.createReadStream("./input"),
	});
	let prev;
	for await (const line of rl) {
		const cur = Number.parseInt(line);

		if (prev == undefined) prev = cur;

		if (cur > prev) count++;

		prev = cur;
	}

	return new Promise<number>((resolve) => resolve(count));
}

(function main() {
	countNumOfIncDepths().then(console.log);
})();
