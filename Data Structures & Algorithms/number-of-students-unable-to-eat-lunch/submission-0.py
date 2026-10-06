class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        n = len(students)
        queue = deque(students)

        res = n

        for sandwich in sandwiches:
            count = 0
            while count < len(queue) and queue[0] != sandwich:
                curr_student = queue.popleft()
                queue.append(curr_student)
                count += 1

            if queue[0] == sandwich:
                queue.popleft()
                res -= 1
            else:
                break

        return res
