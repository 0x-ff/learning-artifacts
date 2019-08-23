
pub mod filewrite {
    use std::io::prelude::*;
    use std::fs::{OpenOptions};

    #[test]
    fn test_write_file() {
        let file_path = "./test_write_file.csv";
        let file_content = "hello write file";
        let mut write_options = OpenOptions::new();
        write_options
            .write(true)
            .create(true);
        match write_options.open(file_path) {
            Ok(mut file) => {
                match file.write_all(file_content.as_bytes()) {
                    Ok(_) => {
                    },
                    Err(err) => {
                        println!("Write file error {:?}\n", err);
                        assert!(false);
                    }
                }
            },
            Err(err) => {
                println!("Create file error {:?}\n", err);
                assert!(false);
            }
        }

        let mut read_options = OpenOptions::new();
        read_options.read(true);
        match read_options.open(file_path) {
            Ok(mut file) => {
                let mut contents = String::new();
                match file.read_to_string(&mut contents) {
                    Ok(_) => {
                        assert_eq!(contents, file_content);
                    },
                    Err(err) => {
                        println!("Read file error {:?}\n", err);
                        assert!(false);
                    }
                }
            },
            Err(err) => {
                println!("Open file error {:?}\n", err);
                assert!(false);
            }
        }
    }
}