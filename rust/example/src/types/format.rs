
pub mod format {

    #[test]
    fn test_string_format() {
        let s = "1234567890";

        assert_eq!(format!("{}", s), s);
        assert_eq!(format!("{:11}", s), s.to_owned() + " ");
        assert_eq!(format!("{:.5}", s), "12345");
        assert_eq!(format!("{:10.5}", s), "12345     ");
        assert_eq!(format!("{:*>10.5}", s), "*****12345");
        assert_eq!(format!("{:*^10.5}", s), "**12345***");
    }

    #[test]
    fn test_number_format() {
        let n = 159;

        assert_eq!(format!("{:+4}", n), "+159");
        assert_eq!(format!("{:+<4}", n), "159+");
        assert_eq!(format!("{:b}", n), "10011111");
    }
}