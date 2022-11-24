use serde::Deserialize;
use walkdir::WalkDir;
use serde_yaml;

fn main() {
    for entry in WalkDir::new(".").into_iter().filter_map(|e| e.ok()) {
        if entry.path().display().to_string().ends_with(".yml") || entry.path().display().to_string().ends_with(".yaml") {
            println!("{}", entry.path().display());
            let f = match std::fs::File::open(entry.path()) {
                Err(e) => {
                    println!("{}", e);
                    continue;
                }
                Ok(f) => f,
            };
            for document in serde_yaml::Deserializer::from_reader(f) {
                let _ :serde_yaml::Value = match serde_yaml::Value::deserialize(document) {
                    Err(e) => {
                        println!("{}", e);
                        continue;
                    }
                    Ok(s) => s,
                };
            }
        }
    }
}