const fs = require('fs');
const csv = require('csv-parser');

const inputCSVFile = 'tasks.csv'; // Replace with your CSV file name
const outputDirectory = 'badges'; // Replace with your desired output directory

// Create the output directory if it doesn't exist
if (!fs.existsSync(outputDirectory)) {
  fs.mkdirSync(outputDirectory);
}

fs.createReadStream(inputCSVFile)
  .pipe(csv())
  .on('data', (row) => {
    const taskName = row.TaskName;
    const status = row.Status;

    // Define badge URL using shields.io service
    const badgeUrl = `https://img.shields.io/badge/${encodeURIComponent(taskName)}-${encodeURIComponent(status)}-${getStatusColor(status)}.svg`;

    // Download the badge image and save it
    const outputPath = `${outputDirectory}/${taskName}.svg`;
    downloadBadgeImage(badgeUrl, outputPath);

    console.log(`Generated badge for ${taskName}`);
  })
  .on('end', () => {
    console.log('Badge generation complete');
  });

// Define a function to map statuses to colors
function getStatusColor(status) {
  switch (status.toLowerCase()) {
    case 'passed':
      return '4CAF50'; // Green
    case 'failed':
      return 'FF5722'; // Red
    case 'skipped':
      return '2196F3'; // Blue
    default:
      return '888888'; // Gray
  }
}

// Function to download and save badge image
const https = require('https');
function downloadBadgeImage(badgeUrl, outputPath) {
  https.get(badgeUrl, (response) => {
    if (response.statusCode === 200) {
      let data = '';
      response.setEncoding('binary');

      response.on('data', (chunk) => {
        data += chunk;
      });

      response.on('end', () => {
        fs.writeFileSync(outputPath, data, 'binary');
      });
    } else {
      console.error(`Failed to download badge from ${badgeUrl}`);
    }
  });
}
