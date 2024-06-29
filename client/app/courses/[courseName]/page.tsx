import React from "react";
import { notFound } from "next/navigation";

import Container from "@/components/common/Container";
import Course from "@/components/courses/Course";
import { HeroHighlight } from "@/components/ui/hero-highlight";
import { getCourse } from "@/actions/course";
import { CourseURLMapper } from "@/config/courses";

type Params = {
  params: {
    courseName: string;
  };
};
export default async function CourseSpecificPage({ params }: Params) {
  const { courseName } = params;

  if (!(courseName in CourseURLMapper)) {
    notFound();
  }
  
  const course = await getCourse(courseName);

  const spacedCourseName = CourseURLMapper[courseName];

  const courseDetails = {
    title: spacedCourseName,
    description: course.description,
    creator: course.creator,
    creatorProfileUrl: "/abyan-150x150.png",
    modulesURL: `/courses/${courseName}`,
    likes: course.likes,
  };

  return (
    <HeroHighlight>
      <Container className="container mx-auto max-w-full px-6 flex-grow">
        <Container className="flex items-start justify-center gap-4 max-w-8xl">
          <Container className="col-span-2">
            <Course courseDetails={courseDetails} />
          </Container>
        </Container>
      </Container>
    </HeroHighlight>
  );
}
