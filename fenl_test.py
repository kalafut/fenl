from unittest import TestCase, main
from fenl import encode_flags

class EncodingTestCase(TestCase):
    def test_encode_flags(self):
        self.assertEqual(encode_flags("w", "KQkq", "-"), "`")
        self.assertEqual(encode_flags("b", "Kq", "-"), "J")
        self.assertEqual(encode_flags("b", "-", "h3"), "Ah")

if __name__ == '__main__':
    main()
