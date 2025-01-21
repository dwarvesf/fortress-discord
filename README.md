# Fortress Discord Bot

Fortress is a Discord bot designed to enhance server interactions by providing various features and functionalities. It integrates with the Discord API to manage messages, commands, and user interactions seamlessly.

## Features
- Health check endpoint to monitor the bot's status.
- Message history management with garbage collection.
- Modular architecture for easy extensibility.

## Getting Started

### Prerequisites
- Go 1.16 or later
- A Discord account to create a bot

### Steps to Set Up

1. Create a bot in Discord.
2. Add the bot to your server with admin permissions and message content intent.
3. Copy the `.env.example` file to `.env`.
4. Fill in the `.env` file with your bot token.
5. Start the bot:

   ```bash
   make dev
   ```

## Directory Structure

- `cmd/`: Contains the entry point for the application.
- `pkg/`: Contains the core functionality, including:
  - `adapter/`: Adapters for various services.
  - `discord/`: Discord-related functionalities, including services and views.
  - `config/`: Configuration management.
  - `logger/`: Logging utilities.
- `.env.example`: Example environment configuration file.
- `Makefile`: Build and run commands.
