/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: "http://localhost:4000/api/v1/:path*",
        // has: [
        //   {
        //     type: "cookie",
        //     key: "authorized",
        //     value: "true",
        //   },
        // ],
      },
    ];
  },
};
