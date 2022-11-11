/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  experimental:{appDir: true},
  async redirects(){
    return[
      {
        source: '/admin',
        destination: '/admin/login',
        permanent: false,
    },
    ]
  },
}
module.exports = nextConfig
