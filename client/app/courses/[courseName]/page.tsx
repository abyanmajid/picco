import React from "react";
import { notFound } from "next/navigation";

import Container from "@/components/common/Container"
import Course from "@/components/courses/Course";
import { courses, CourseURLMapper } from "@/config/courses"
import { HeroHighlight } from "@/components/ui/hero-highlight";

type Params = {
    params: {
        courseName: string
    }
}
export default function CourseSpecificPage({ params }: Params) {
    const { courseName } = params;

    if (!(courseName in CourseURLMapper)) {
        notFound();
    }

    const spacedCourseName = CourseURLMapper[courseName]

    const courseDetails = {
        title: spacedCourseName,
        description: courses[spacedCourseName].description,
        creator: courses[spacedCourseName].creator,
        creatorProfileUrl: courses[spacedCourseName].creatorProfileUrl,
        courseUrl: courses[spacedCourseName].courseUrl,
        likes: courses[spacedCourseName].likes,
    }

    return (
        <HeroHighlight>
            <Container className="container mx-auto max-w-full px-6 flex-grow">
                <Container className="flex items-start justify-center gap-4 max-w-8xl">
                    <Container className="col-span-2">
                        <Course courseDetails={courseDetails} />
                    </Container>
                </Container>
            </Container>
        </HeroHighlight >
    );
}