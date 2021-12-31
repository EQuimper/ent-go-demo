import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Input,
  Text,
} from "@chakra-ui/react";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { useQueryClient } from "react-query";
import { useAuth } from "../components/hooks/useAuth";

function Register() {
  const queryClient = useQueryClient();
  const router = useRouter();
  // const { isLogged } = useAuth();

  // useEffect(() => {
  //   if (isLogged) {
  //     router.push("/projects");
  //   }
  // }, []);

  const onSubmit = (e: any) => {
    e.preventDefault();

    const { email, password, password_confirmation, username } =
      e.target.elements;

    fetch("/api/auth/register", {
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        password_confirmation: password_confirmation.value,
        username: username.value,
      }),
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      method: "POST",
      credentials: "include",
    })
      .then((res) => res.json())
      .then(() => {
        queryClient.invalidateQueries("me");
        router.push("/projects");
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <Box mt={8}>
      <Container>
        <Text fontSize="2xl" fontWeight="bold" mb={8}>
          Register
        </Text>
        <form onSubmit={onSubmit}>
          <FormControl isRequired>
            <FormLabel id="username" htmlFor="username">
              Username
            </FormLabel>
            <Input
              id="username"
              name="username"
              placeholder="Username"
              required
            />
          </FormControl>

          <FormControl isRequired mt={4}>
            <FormLabel id="email" htmlFor="email">
              Email
            </FormLabel>
            <Input
              id="email"
              name="email"
              placeholder="Email"
              type="email"
              required
            />
          </FormControl>

          <FormControl isRequired mt={4}>
            <FormLabel id="password" htmlFor="password">
              Password
            </FormLabel>
            <Input
              id="password"
              name="password"
              placeholder="Password"
              type="password"
              required
            />
          </FormControl>

          <FormControl isRequired mt={4}>
            <FormLabel
              id="password_confirmation"
              htmlFor="password_confirmation"
            >
              Password Confirmation
            </FormLabel>
            <Input
              id="password_confirmation"
              name="password_confirmation"
              placeholder="Password confirmation"
              type="password"
              required
            />
          </FormControl>

          <Button mt={8} colorScheme="blue" w="100%" type="submit">
            Submit
          </Button>
        </form>
      </Container>
    </Box>
  );
}

export function getServerSideProps({ req }) {
  if (req.cookies.authorization) {
    return {
      redirect: {
        destination: "/projects",
        permanent: false,
      },
    };
  }

  return {
    props: {},
  };
}

export default Register;
