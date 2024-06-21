import CourseCard from "./CourseCard";
import Container from "../common/Container";

export default function HotCourses() {
    return <Container className="grid grid-cols-2 gap-4 mt-4 py-4">
        <CourseCard
            title="Introduction to Programming"
            description="Learn fundamental programming constructs, problem-solving, and object-oriented design in python."
            modules={26}
            courseUrl="/courses/introduction-to-programming"
        />
        <CourseCard
            title="Data Structures and Algorithms"
            description="Learn asymptotic analysis, data structures, important algorithms, and correctness proofing."
            modules={14}
            courseUrl="/courses/data-structures-and-algorithms"
        />
    </Container>
}