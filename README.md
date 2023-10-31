# RSS Aggregate Service in Golang

![RSS Aggregate Service](rssagg.png)

This is a simple RSS Aggregate Service written in Golang that allows users to subscribe to RSS feeds and have those feeds updated automatically. It's a lightweight and easy-to-use service for aggregating and staying up-to-date with your favorite websites and blogs.

## Features

- Subscribe to multiple RSS feeds.
- Automatic feed updates at regular intervals.
- Easily manage and organize your subscribed feeds.
- Simple and intuitive web-based user interface.
- Customizable update frequency.
- Access your aggregated feeds from anywhere.

## Prerequisites

Before you can run the RSS Aggregate Service, you need to have the following prerequisites installed on your system:

- [Golang](https://golang.org/) (at least Go 1.16)
- [Git](https://git-scm.com/)

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/rssagg.git
   ```

2. Change into the project directory:

   ```bash
   cd rssagg
   ```

3. Build the project:

   ```bash
   go build
   ```

4. Run the RSS Aggregate Service:

   ```bash
   ./rssagg
   ```

The service should now be running on `http://localhost:8080`. You can access it through your web browser.

## Usage

1. Access the web interface by opening a web browser and navigating to `http://localhost:8080`.

2. Create an account or log in if you already have one.

3. Once logged in, you can start adding RSS feeds to your account.

4. Click on the "Add Feed" button and enter the URL of the RSS feed you want to subscribe to.

5. Choose the update frequency for the feed (e.g., every hour, every day).

6. Click "Subscribe" to add the feed to your list.

7. The service will automatically update the subscribed feeds at the specified intervals.

8. You can view and read the aggregated feeds on the dashboard.

9. To manage your subscribed feeds, click on the "My Feeds" or "Settings" sections.

## Configuration

You can customize the service by modifying the `config.yaml` file. Here are some of the configuration options:

- `port`: The port on which the service will run.
- `update_interval`: The default update interval for newly subscribed feeds.
- `max_feeds_per_user`: The maximum number of feeds a user can subscribe to.
- `database_path`: The path to the SQLite database file.

## Contributing

If you would like to contribute to the development of this project, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them thoroughly.
4. Commit your changes and push them to your forked repository.
5. Submit a pull request to the main repository, explaining your changes and why they should be merged.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project was inspired by the need for a simple and self-hosted RSS aggregation service.
- Thanks to the Golang community for creating a powerful and efficient programming language.

## Contact

If you have any questions or feedback, please feel free to reach out to us at [dmmoody@gmail.com](mailto:your.email@example.com).

Happy RSS feed aggregating! 🚀
