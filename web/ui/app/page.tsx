import Hero from "@/components/landing/Hero";
import HotCourses from "@/components/landing/HotCourses";
import Container from "@/components/common/Container";

export default function HomePage() {
  return (
    <Container className="flex flex-col items-center justify-center gap-4 mt-20 pt-16">
      <Hero />
      <HotCourses />
    </Container>
  );
}
