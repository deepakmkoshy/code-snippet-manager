
// const { default: axios } = require('axios');
const vscode = require('vscode');
const { dirname } = require('path');
const fs = require('fs');
const path = require('path');

/**
 * @param {vscode.ExtensionContext} context
 */
function activate(context) {

	console.log('Congratulations, your extension "code-snippet-manager" is now active!');

	let disposable = vscode.commands.registerCommand('code-snippet-manager.createCodeSnippet', async function () {
		// Get Selected Text
		let editor = vscode.window.activeTextEditor;
		if (editor) {
			const selection = editor.selection;
			const text = editor.document.getText(selection);
			const selectedChunk = text.replace("\r", '').split("\n");
			const body = selectedChunk.length === 1 ? selectedChunk[0] : selectedChunk;

			//Ask user for prefix
			const prefix = await vscode.window.showInputBox({
				prompt: "Enter prefix for snippet",
				placeHolder: "prefix"
			});

			const snippet = {
				"Short Description": {
					prefix,
					body,
					description: "Long Description"
				}
			}


			const snippetString = JSON.stringify(snippet);
			// fs.mkdirSync(path.join(__dirname, '.vddscode'), { recursive: true });

			fs.appendFileSync(path.join(__dirname, '.vscode'), snippetString, function (err) {
				if (err) throw err;
				console.log('Saved!');
			});

			// // remove first and last char from string
			// const snippetStringWithoutBrackets = snippetString.substring(1, snippetString.length - 1);
			// // remove leading and trailing newline from a string
			// const snippetStringWithoutNewline = snippetStringWithoutBrackets.replace(/^\s*[\r\n]/gm, '').replace(/[\r\n]\s*$/gm, '');
			// // Copy to clipboard
			// vscode.env.clipboard.writeText(snippetStringWithoutNewline.trim());

			// console.log(snippetStringWithoutNewline.trim());
			// // Show copied Message
			// vscode.window.showInformationMessage('Snippet Copied to clipboard 📋');

		}
	});

	context.subscriptions.push(disposable);
}


function deactivate() { }

module.exports = {
	activate,
	deactivate
}
