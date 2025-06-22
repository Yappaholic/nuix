use super::commands::SystemCmd;
pub fn run(command: &SystemCmd) {
    if let Some(cmd) = command.command {
        println!("Called with command {:?}", cmd);
    }
    if let Some(flake_path) = &command.flake {
        println!("Called with flake path {:?}", flake_path);
    }
    if command.nh == true {
        println!("Will use nh for calls");
    }
}
