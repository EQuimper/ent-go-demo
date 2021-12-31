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
import { useQueryClient } from "react-query";

function Login() {
  const router = useRouter();
  const queryClient = useQueryClient();
  const onSubmit = (e: any) => {
    e.preventDefault();

    const { email, password } = e.target.elements;

    fetch("/api/auth/login", {
      body: JSON.stringify({
        email: email.value,
        password: password.value,
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
        router.push("/projects");
        queryClient.invalidateQueries("me");
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <Box mt={8}>
      <Container>
        <Text fontSize="2xl" fontWeight="bold" mb={8}>
          Login
        </Text>
        <form onSubmit={onSubmit}>
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

          <Button mt={8} colorScheme="blue" w="100%" type="submit">
            Login
          </Button>
        </form>
      </Container>
    </Box>
  );
}

export default Login;
