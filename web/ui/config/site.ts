export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  name: "codemore.io",
  description: "Learn at the speed of light by simply writing more code.",
  navItems: [
    {
      label: "Home",
      href: "/",
    },
    {
      label: "Courses",
      href: "/courses",
    },
    {
      label: "Top Learners",
      href: "/leaderboard",
    },
  ],
  navMenuItems: [
    {
      label: "Home",
      href: "/",
    },
    {
      label: "Courses",
      href: "/courses",
    },
    {
      label: "Top Learners",
      href: "/leaderboard",
    },
  ],
  links: {
    github: "https://github.com/abyanmajid/codemore.io",
    login: "/api/auth/login",
    logout: "/api/auth/logout",
  },
};
