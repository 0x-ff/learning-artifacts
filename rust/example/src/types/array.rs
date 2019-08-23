
pub mod array {

    #[test]
    fn test_len() {
        let arr: [u32; 6] = [1, 2, 4, 7, 11, 16];
        assert_eq!(6, arr.len());
    }

    #[test]
    fn test_vec_reverse() {
        let mut v = vec!["a", "b", "c"];
        v.reverse();
        assert_eq!(v, vec!["c", "b", "a"]);
    }

    #[test]
    fn test_vec_pop() {
        let mut v = vec!["a", "b", "c"];
        assert_eq!(v.pop(), Some("c"));
        assert_eq!(v.pop(), Some("b"));
        assert_eq!(v.pop(), Some("a"));
        assert_eq!(v.pop(), None);
        assert_eq!(v.pop(), None);
    }

    #[test]
    fn test_slices() {
        let v = vec!["a", "b", "c", "d", "e"];
        assert_eq!(v[0..2], ["a", "b"]);
        assert_eq!(v[2..], ["c", "d", "e"]);
        assert_eq!(v[1..3], ["b", "c"]);
        assert_eq!(v[..3], ["a", "b", "c"]);
    }
}