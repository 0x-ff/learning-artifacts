

pub mod numbers {
    use panic_control::{Context, spawn_quiet, ThreadResultExt};

    #[derive(Debug, PartialEq, Eq)]
    #[allow(dead_code)]
    enum Expected {
         Token,
         Int(i32),
         String(String)
    }

    #[test]
    fn test_u8_overflow() {
        let ctx = Context::<Expected>::new();
        let h = ctx.spawn_quiet(|| {
            let h = spawn_quiet(|| {
                let mut value: u8;
                value = 255;
                value = value + 1;
                assert_eq!(0, value);
            });
            h.join().unwrap_or_propagate();
        });
        let result = h.join();
        let message = result.panic_value_as_str().unwrap();
        assert!(message.contains("attempt to add with overflow"));
    }

    #[test]
    fn test_i8_overflow() {
        let ctx = Context::<Expected>::new();
        let h = ctx.spawn_quiet(|| {
            let h = spawn_quiet(|| {
                let mut value: i8;
                value = -128;
                value = value - 1;
                assert_eq!(0, value);
            });
            h.join().unwrap_or_propagate();
        });
        let result = h.join();
        let message = result.panic_value_as_str().unwrap();
        assert!(message.contains("attempt to subtract with overflow"));
    }

    #[test]
    fn test_f32_overflow() {
        let mut value: f32;
        value = 3.4_e38;
        value = value * 2.0;

        assert_eq!(std::f32::INFINITY, value);
    }


    #[test]
    fn test_f64_overflow() {
        let mut value: f64;
        value = -1.7_e308;
        value = value * 2.0;

        assert_eq!(std::f64::NEG_INFINITY, value);
    }
}