import { Box, Container, List, ListItem, Text } from "@chakra-ui/react";
import { useQuery } from "react-query";
import * as cookie from "cookie";

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

function Projects(props) {
  const { data, isLoading } = useQuery("projects", getProjects, {
    initialData: props.projects,
  });

  if (isLoading) {
    return null;
  }

  // console.log("data client", data);
  return (
    <Box bg="gray.50" minHeight="100vh" pt={8}>
      <Container>
        <Text fontSize="2xl" fontWeight="bold">
          My Projects
        </Text>

        <List>
          {data.map((project) => (
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
  console.log("data server", data);

  return {
    props: {
      projects: data,
    },
  };
}

export default Projects;
