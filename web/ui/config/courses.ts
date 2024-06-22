type Course = {
    title: string,
    description: string,
    creator: string,
    creatorProfileUrl: string,
    courseUrl: string,
    likes: number,
    modules: {
        name: string,
        type: "lecture" | "exercise" | "quiz" | "guided project",
        xp: number,
    }[]
}

type CourseMap = {
    [key: string]: string;
};

export const courses: Record<string, Course> = {
    "Introduction to Programming": {
        title: "Introduction to Programming",
        description: "Learn fundamental programming constructs, problem-solving, and object-oriented design in python.",
        creator: "Abyan Majid",
        creatorProfileUrl: "https://github.com/abyanmajid",
        courseUrl: "/introduction-to-programming",
        likes: 107,
        modules: [
            {
                name: "Variables and data types",
                type: "lecture",
                xp: 150,
            },
            {
                name: "Leap year",
                type: "exercise",
                xp: 320,
            },
            {
                name: "Conditionals Quiz",
                type: "quiz",
                xp: 170,
            },
        ]
    }
}

export const CourseURLMapper: CourseMap = {
    "introduction-to-programming": "Introduction to Programming",
    "data-structures-and-algorithms": "Data Structures and Algorithms"
}