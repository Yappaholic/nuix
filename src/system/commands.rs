use std::path::PathBuf;

use clap::{Parser, Subcommand};

#[derive(Parser, Debug)]
pub struct SystemCmd {
    #[command(subcommand)]
    pub command: Option<SystemArg>,

    /// use flake for configuration
    #[arg(short, long, value_name = "PATH")]
    pub flake: Option<PathBuf>,

    /// prefer nh for commands
    #[arg(long)]
    pub nh: bool,
}

#[derive(Subcommand, Debug, Clone, Copy)]
pub enum SystemArg {
    /// switch to the next generation
    Switch,
    /// test new configuration without switching
    Test,
    /// build new configuration and do nothing after
    Build,
    /// build new configuration but only show changes
    DryTest,
    /// say moo
    Moo,
}
