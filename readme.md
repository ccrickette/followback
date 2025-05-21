# Followback for Instagram

Followback is a simple and privacy-focused tool for Instagram users who want to see which of their followers are not following them back. Unlike third-party apps that might access your private data, Followback ensures your data stays secure by only using the follower information you download from Instagram yourself.

## What is Followback?

This program allows you to easily compare your followers and the people you follow on Instagram to identify users who do not follow you back. It helps you manage your Instagram connections without requiring access to your private data or relying on third-party services that might compromise your account security.

## Why was Followback made?

Instagram users often want to know who isn't following them back, but most apps that offer this service require access to sensitive data. These third-party apps could potentially steal your information. Followback eliminates that risk by letting you download your own follower information directly from Instagram and analyze it locally on your machine.

## How to Use Followback

1. **Download Your Information from Instagram**
   - Visit [Instagram's Data Download page](https://accountscenter.instagram.com/info_and_permissions/)
   - Select "Your information and permissions"
   - Choose "Download your information"
   - Select "Some of your information" (we only need your followers list, nothing else)
   - Under the "Connections" header, choose "Followers and Following"
   - Make sure to select "All Time" to get the full list of your followers and following history.
   - Make sure you select media quality Medium, and format JSON!
   - Click "Request Download." It may take some time for Meta (Instagram's parent company) to prepare your data and will be in progess.

2. **Download and Extract the Zip File**
   - Once Instagram has processed the request, you will receive a zip file containing two JSON files: `followers_1.json` and `following.json`.
   - Download and extract the zip file.

3. **Run the Program**
   - Place the two JSON files (`followers_1.json` and `following.json`) in the root directory of this project.
   - Run the program, and it will generate a `difference.txt` file containing a list of people who do not follow you back.

4. **View Results**
   - Open `difference.txt` to see the list of users who are not following you back on Instagram.

## Example Output

The `difference.txt` file will contain a simple list of Instagram usernames who are following you but are not following you back.

---

## Why is This Safe to Use?

Followback only works with the follower and following data you download directly from Instagram. We don’t store or process your data on any external server. All the data remains local to your machine, and the program simply compares the two lists to identify discrepancies.

## Prerequisites

- Go compiler
- Access to your Instagram follower data (as described above)

## License

This project is open-source and available under the MIT License. See the LICENSE file for more details.
