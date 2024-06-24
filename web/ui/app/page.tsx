import Hero from "@/components/landing/Hero";
import HotCourses from "@/components/landing/HotCourses";
import Container from "@/components/common/Container";
import { HeroHighlight } from "@/components/ui/hero-highlight";
import { getSession, withPageAuthRequired } from '@auth0/nextjs-auth0';

export default async function HomePage() {

  return (
    <HeroHighlight>
      <Container className="container mx-auto max-w-full px-6 flex-grow">
        <Container className="flex flex-col items-center justify-center gap-4 mt-20 pt-16">
          <Hero />
          <HotCourses />
        </Container>
      </Container>
    </HeroHighlight>
  );
}
