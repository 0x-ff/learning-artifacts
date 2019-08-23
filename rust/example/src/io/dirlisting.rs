
pub mod dirlisting {
    use std::path::Path;
    use std::io::ErrorKind;

    #[test]
    fn test_dir_listing() {
        let mut found = false;

        match Path::new(".").read_dir() {
            Ok(result) => {
                for entry_result in result {
                    match entry_result {
                        Ok(entry) => {
                            if entry.file_name().to_string_lossy() == "Cargo.lock" {
                                found = true;
                                break;
                            }
                        },
                        Err(err) => {
                            println!("Entry result error {:?}\n", err);
                            assert!(false);
                        }
                    };
                }
            },
            Err(err) => {
                println!("Read dir error {:?}\n", err);
                assert!(false);
            }
        };
        assert_eq!(true, found);
    }

    #[test]
    fn test_dir_not_found() {
        match Path::new("<>").read_dir() {
            Ok(_) => {
                assert!(false);
            },
            Err(err) => {
                assert_eq!(ErrorKind::NotFound, err.kind());
            }
        }
    }
}
