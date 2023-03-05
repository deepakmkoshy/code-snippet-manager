const fs = require('fs');


const snippet = {
    "Short Description": {
        prefix: "prefix",
        body: 'body',
        description: "Long Description"
    }
}

const snippetString = JSON.stringify(snippet);

fs.appendFileSync('.vscode/snippet.code-snippets', snippetString, function (err) {
    if (err) throw err;
    console.log('Saved!');
});
