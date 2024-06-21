import React from "react";

import Container from "@/components/common/Container"
import Title from "@/components/common/Title";
import Subtitle from "@/components/common/Subtitle";
import CourseCard from "@/components/landing/CourseCard";

export default function CoursesPage() {
    return (
        <Container className="flex flex-col justify-center items-center space-y-4">
            <Title size="sm">Available Courses</Title>
            <Subtitle className="text-center">Dive in, get hands-on, and experience rapid growth.</Subtitle>
            <Container className="grid grid-cols-2 gap-4">
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

        </Container>
    );
}