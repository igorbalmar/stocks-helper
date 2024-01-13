/* global use, db */
// MongoDB Playground
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.

const database = 'stocks-helper';
const collection = 'stock-prices';

// The current database to use.
use(database);

// The prototype form to create a collection:
db.createCollection( "stock-prices", { timeseries: { timeField: "timestamp"} })


// More information on the `createCollection` command can be found at:
// https://www.mongodb.com/docs/manual/reference/method/db.createCollection/
