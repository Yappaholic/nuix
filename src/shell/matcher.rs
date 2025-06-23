use super::commands::Commands;

pub fn run(command: &Commands) {
    let packages = command.packages.clone().take().unwrap();
    let packages = packages.split(',');
    println!(
        "will call nix-shell with flake: {:?}",
        command.flake.clone().take(),
    );
    print!("will use next packages: ");
    for n in packages {
        print!("{} ", n);
    }
}
