const fs = require('fs-extra')

async function copyFiles () {
    await fs.remove('../../public')
    await fs.copy('./build', '../../public')
  }

copyFiles().then(c => {
    console.log("done.")
}).catch(err => {
    console.error(err);
    process.exit(1);
});