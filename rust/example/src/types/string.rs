

pub mod string {

    #[test]
    fn test_trim() {
        assert_eq!("\tab c \n".trim(), "ab c");
        assert_eq!("ab d".trim(), "ab d");

        assert_eq!("\tab c \n".trim_start(), "ab c \n");
        assert_eq!("ab d".trim_start(), "ab d");

        assert_eq!("\tab c \n".trim_end(), "\tab c");
        assert_eq!("ab d".trim_end(), "ab d");

        assert_eq!("ccab cc".trim_matches('c'), "ab ");
        assert_eq!("ab d ".trim_matches('c'), "ab d ");

        assert_eq!("ccab cc".trim_start_matches('c'), "ab cc");
        assert_eq!("ab d ".trim_start_matches('c'), "ab d ");

        assert_eq!("ccab cc".trim_end_matches('c'), "ccab ");
        assert_eq!("ab d ".trim_end_matches('c'), "ab d ");
    }

    #[test]
    fn test_split() {
        let mut iter = "ab;;cd;;aa;a;;".split(";;");
        assert_eq!(iter.next(), Some("ab"));
        assert_eq!(iter.next(), Some("cd"));
        assert_eq!(iter.next(), Some("aa;a"));
        assert_eq!(iter.next(), Some(""));
        assert_eq!(iter.next(), None);

        iter = "ab;;cd;;aa;a".split("::");
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_splitn() {
        let mut iter = "ab;;cd;;aa;a;;".splitn(2, ";;");
        assert_eq!(iter.next(), Some("ab"));
        assert_eq!(iter.next(), Some("cd;;aa;a;;"));
        assert_eq!(iter.next(), None);

        iter = "ab;;cd;;aa;a".splitn(2, "::");
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_rsplit() {
        let mut iter = "ab;;cd;;aa;a;;".rsplit(";;");
        assert_eq!(iter.next(), Some(""));
        assert_eq!(iter.next(), Some("aa;a"));
        assert_eq!(iter.next(), Some("cd"));
        assert_eq!(iter.next(), Some("ab"));
        assert_eq!(iter.next(), None);

        iter = "ab;;cd;;aa;a".rsplit("::");
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_rsplitn() {
        let mut iter = "ab;;cd;;aa;a;;".rsplitn(2, ";;");
        assert_eq!(iter.next(), Some(""));
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);

        iter = "ab;;cd;;aa;a".rsplitn(2, "::");
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_split_terminator() {
        let mut iter = ";;ab;;cd;;aa;a;;".split_terminator(";;");
        assert_eq!(iter.next(), Some(""));
        assert_eq!(iter.next(), Some("ab"));
        assert_eq!(iter.next(), Some("cd"));
        assert_eq!(iter.next(), Some("aa;a"));
        assert_eq!(iter.next(), None);

        iter = "ab;;cd;;aa;a".split_terminator("::");
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_rsplit_terminator() {
        let mut iter = ";;ab;;cd;;aa;a;;".rsplit_terminator(";;");
        assert_eq!(iter.next(), Some("aa;a"));
        assert_eq!(iter.next(), Some("cd"));
        assert_eq!(iter.next(), Some("ab"));
        assert_eq!(iter.next(), Some(""));
        assert_eq!(iter.next(), None);

        iter = "ab;;cd;;aa;a".rsplit_terminator("::");
        assert_eq!(iter.next(), Some("ab;;cd;;aa;a"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_matches() {
        let mut iter = ";;ab;;cd;;aa;a;;".matches(";;");
        assert_eq!(iter.next(), Some(";;"));
        assert_eq!(iter.next(), Some(";;"));
        assert_eq!(iter.next(), Some(";;"));
        assert_eq!(iter.next(), Some(";;"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_match_indices() {
        let mut iter = ";;ab;;cd;;aa;a;;".match_indices(";;");
        assert_eq!(iter.next(), Some((0, ";;")));
        assert_eq!(iter.next(), Some((4, ";;")));
        assert_eq!(iter.next(), Some((8, ";;")));
        assert_eq!(iter.next(), Some((14, ";;")));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_contains() {
        assert!(";;ab;;cd;;aa;a;;".contains(";;cd;"));
        assert!(!";;ab;;cd;;aa;a;;".contains(";;cd;;;"));
    }

    #[test]
    fn test_starts_with() {
        assert!(";;ab;;cd;;aa;a;;".starts_with(";;ab;"));
        assert!(!";;ab;;cd;;aa;a;;".starts_with(";;ab;;;"));
    }

    #[test]
    fn test_ends_with() {
        assert!(";;ab;;cd;;aa;a;;".ends_with("a;a;;"));
        assert!(!";;ab;;cd;;aa;a;;".ends_with(";;ab;;;"));
    }

    #[test]
    fn test_find() {
        assert_eq!("abcdefcdefac".find("d"), Some(3));
        assert_eq!("abcdefcdefac".find("z"), None);
    }

    #[test]
    fn test_rfind() {
        assert_eq!("abcdefcdefac".rfind("d"), Some(7));
        assert_eq!("abcdefcdefac".rfind("z"), None);
    }

    #[test]
    fn test_replace() {
        assert_eq!("abcdefcdefac".replace("d", "DcD"), "abcDcDefcDcDefac");
        assert_eq!("abcdefcdefac".replace("z", "f"), "abcdefcdefac");
    }

    #[test]
    fn test_replacen() {
        assert_eq!("abcdefcdefac".replacen("d", "DcD", 1), "abcDcDefcdefac");
        assert_eq!("abcdefcdefac".replacen("z", "f", 1), "abcdefcdefac");
    }

    #[test]
    fn test_drain() {
        let mut s = "abcDefcdefac".to_string(); 
        assert_eq!(s.drain(3..6).collect::<String>(), "Def");
        assert_eq!(s, "abccdefac");
    }

    #[test]
    fn test_insert() {
        let mut s = "1234567890".to_string(); 
        s.insert(2, 'a');
        assert_eq!(s, "12a34567890");

        s.insert(0, 'b');
        assert_eq!(s, "b12a34567890");

        s.insert(s.len(), 'c');
        assert_eq!(s, "b12a34567890c");
    }
}