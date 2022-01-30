import { Box, Container, List, ListItem, Text } from "@chakra-ui/react";
import { useQuery } from "react-query";

interface Project {
  id: number;
  name: string;
  description?: string;
  created_at: string;
}

async function getProjects() {
  const res = await fetch("/api/projects", {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
  });
  const { data } = await res.json();

  return data;
}

interface Props {
  projects?: Project[];
}

function Projects(props: Props) {
  const { data, isLoading } = useQuery<Project[]>("projects", getProjects, {
    initialData: props.projects,
  });

  if (isLoading) {
    return null;
  }

  const projects = data ?? [];

  return (
    <Box bg="gray.50" minHeight="100vh" pt={8}>
      <Container>
        <Text fontSize="2xl" fontWeight="bold">
          My Projects
        </Text>

        <List>
          {projects.map((project) => (
            <ListItem
              key={project.id}
              boxShadow="base"
              borderRadius="lg"
              mt={4}
            >
              <Box p={6}>
                <Text fontSize="lg" fontWeight="medium">
                  {project.name}
                </Text>

                <Text color="gray.300">{project.description ?? ""}</Text>
              </Box>
            </ListItem>
          ))}
        </List>
      </Container>
    </Box>
  );
}

export async function getServerSideProps({ req }) {
  if (!req.cookies.authorization) {
    return {
      redirect: {
        destination: "/",
        permanent: false,
      },
    };
  }

  const response = await fetch("http://localhost:4000/api/v1/projects", {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      Cookie: req.headers.cookie,
    },
  });
  const { data } = await response.json();

  return {
    props: {
      projects: data,
    },
  };
}

export default Projects;
