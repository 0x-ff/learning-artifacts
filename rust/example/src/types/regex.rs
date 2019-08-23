

pub mod regex {
    use regex::Regex;

    #[test]
    fn test_is_match() {
        match Regex::new(r"(\d{4})\-(\d{2})-(\d{2})") {
            Ok(pattern) => {
                assert!(pattern.is_match("2019-01-01"));
                assert!(!pattern.is_match("2019-01-0a2"));
            },
            Err(err) => {
                println!("{}", err);
                assert!(false);
            }
        };
    }

    #[test]
    fn test_captures() {
        match Regex::new(r"(\d{4})\-(\d{2})-(\d{2})") {
            Ok(pattern) => {
                match pattern.captures("2019-02-01") {
                    Some(matches) => {
                        assert_eq!(&matches[0], "2019-02-01");
                        assert_eq!(&matches[1], "2019");
                        assert_eq!(&matches[2], "02");
                        assert_eq!(&matches[3], "01");
                    },
                    None => {
                        println!("No matches!");
                        assert!(false);
                    }
                };
            },
            Err(err) => {
                println!("{}", err);
                assert!(false);
            }
        };
    }
}