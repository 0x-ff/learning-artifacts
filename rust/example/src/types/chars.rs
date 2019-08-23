
pub mod chars {
    
    #[test]
    fn test_is_numeric() {
        assert!('4'.is_numeric());
        assert!(!'a'.is_numeric());
    }

    #[test]
    fn test_is_alphabetic() {
        assert!(!'4'.is_alphabetic());
        assert!('a'.is_alphabetic());
        assert!('Ю'.is_alphabetic());
    }

    #[test]
    fn test_is_alphanumeric() {
        assert!('4'.is_alphanumeric());
        assert!('a'.is_alphanumeric());
        assert!('ю'.is_alphanumeric());
        assert!('ё'.is_alphanumeric());
    }

    #[test]
    fn test_is_whitespace() {
        assert!(!'4'.is_whitespace());
        assert!(' '.is_whitespace());
        assert!('\n'.is_whitespace());
        assert!('\u{A0}'.is_whitespace());
        assert!('\r'.is_whitespace());
        assert!('\t'.is_whitespace());
    }

    #[test]
    fn test_is_control() {
        assert!(!'4'.is_control());
        assert!('\n'.is_control());
        assert!('\u{85}'.is_control());
    }

    #[test]
    fn test_is_digit() {
        assert!('4'.is_digit(10));
        assert!(!'A'.is_digit(10));
        assert!('A'.is_digit(16));
    }

    #[test]
    fn test_to_digit() {
        match 'A'.to_digit(16) {
            Some(num) => {
                assert_eq!(num, 10);
            },
            None => {
                assert!(false);
            }
        }
    }

    #[test]
    fn test_is_uppercase() {
        assert!(!'4'.is_uppercase());
        assert!('Z'.is_uppercase());
        assert!(!'z'.is_uppercase());
        assert!('Б'.is_uppercase());
        assert!(!'б'.is_uppercase());
    }

    #[test]
    fn test_is_lowercase() {
        assert!(!'4'.is_lowercase());
        assert!(!'Z'.is_lowercase());
        assert!('z'.is_lowercase());
        assert!(!'Б'.is_lowercase());
        assert!('б'.is_lowercase());
    }

    #[test]
    fn test_to_uppercase() {
        let mut upper = 'a'.to_uppercase();
        assert_eq!(upper.next(), Some('A'));
        assert_eq!(upper.next(), None);

        upper = '1'.to_uppercase();
        assert_eq!(upper.next(), Some('1'));
        assert_eq!(upper.next(), None);
    }

    #[test]
    fn test_to_lowercase() {
        let mut lower = 'A'.to_lowercase();
        assert_eq!(lower.next(), Some('a'));
        assert_eq!(lower.next(), None);

        lower = '1'.to_lowercase();
        assert_eq!(lower.next(), Some('1'));
        assert_eq!(lower.next(), None);
    }
}