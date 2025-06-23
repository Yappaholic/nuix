use std::path::PathBuf;

use clap::Parser;

#[derive(Parser, Debug)]
pub struct Commands {
    /// use flake
    #[arg(short, long, value_name = "path")]
    pub flake: Option<PathBuf>,

    /// add argument
    #[arg(long, value_name = "name value")]
    pub arg: Option<Vec<String>>,

    #[arg(long, value_name = "name value")]
    pub argstr: Option<Vec<String>>,

    #[arg(long, short, value_name = "attrPath")]
    pub attr: Option<PathBuf>,

    #[arg(long, value_name = "cmd")]
    pub command: Option<String>,

    #[arg(long, value_name = "cmd")]
    pub run: Option<String>,

    #[arg(long, value_name = "regexp")]
    pub regexp: Option<String>,

    #[arg(long)]
    pub pure: bool,

    #[arg(short, long, value_name = "packages | expressions", value_parser = check_packages)]
    pub packages: Option<String>,
}

fn check_packages(s: &str) -> Result<String, String> {
    let s_temp: String = s.split_whitespace().collect();
    let s_temp = s_temp.as_bytes();
    println!("{}", s_temp.len());
    for i in 0..s_temp.len() {
        let c: char = s_temp[i] as char;
        println!("{}", c);
        if c == ',' {
            continue;
        }
        if !c.is_alphanumeric() {
            return Err(format!("package list has illegal characters"));
        }
    }
    Ok(s.to_owned())
}
