# wallet-buddy

Your trusted companion for managing all your financial assets in one place

## Database Setup

Update the `database.yml` file with the correct database name, username, and password for your environment. Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

```console
buffalo pop create -a
```

This will create the databases, run all the migrations, and then populate the database with seed data.

## Running Migrations

Run this command to run all migrations:

```console
buffalo pop migrate
```

## Running Seeds

Run this command to run all seeds:

```console
buffalo pop seed
```

## Manage application ENV variables

Create a `.env` file in the root of the project and add the following variables:

```console
PORT=30001 # Port on which the application will run
```

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```console
buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Wallet Buddy application up and running.

To list all routes for the application, run the following command:

```console
buffalo task routes
```

Happy coding!

[Powered by Buffalo](http://gobuffalo.io)
