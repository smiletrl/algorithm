import unittest
from dataclasses import dataclass
from s import *

class TestSolution(unittest.TestCase):
    def test_reverseKGroup(self):
        @dataclass
        class TestCase:
            name: str
            input: list[int]
            key: int
            expect: list[int]
        
        testCases = [
            TestCase(name="case 1", input=[1,4,5,7,8], key=2, expect=[4,1,7,5,8])
        ]

        for ca in testCases:
            head = ListNode()
            ohead = head
            for i in range(len(ca.input)):
                head.val = ca.input[i]
                if i != len(ca.input)-1:
                    node = ListNode()
                    head.next = node
                    head = node
            s = Solution()
            out = s.reverseKGroup(ohead, ca.key)
            x = 0
            while out != None:
                self.assertEqual(ca.expect[x], out.val, "expect is not equal to real")
                x += 1
                out = out.next

if __name__ == "__main__":
     unittest.main()
