mod home;
mod shell;
mod system;
use clap::{Parser, Subcommand};

#[derive(Parser, Debug)]
struct Cli {
    #[command(subcommand)]
    command: Option<Domain>,
}

#[derive(Subcommand, Debug)]
enum Domain {
    System(system::commands::SystemCmd),
    Home(home::commands::HomeCmd),
    Shell(shell::commands::Commands),
}

fn main() {
    let cli = Cli::parse();
    matcher(cli);
}

fn matcher(cli: Cli) {
    match &cli.command {
        Some(Domain::System(name)) => {
            system::matcher::run(name);
        }
        Some(Domain::Home(name)) => {
            home::matcher::run(name);
        }
        Some(Domain::Shell(name)) => {
            shell::matcher::run(name);
        }
        None => {}
    }
}
