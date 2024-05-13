<img src="https://raw.githubusercontent.com/eneru2/wibo/main/.github/media/wibo_long_logo.png" alt="Wibo logo">

# Wibo
Wibo is an advanced image optimization tool and code generator designed to assist developers in seamlessly converting images into the three primary web formats: AVIF, WebP, and JPG. It not only optimizes images for modern web standards but also ensures compatibility across various browsers by automatically generating fallback code. This means that while Wibo prioritizes newer formats for enhanced performance and quality, it intelligently falls back to alternative options when compatibility issues arise, ensuring a consistent user experience across different platforms and devices.

## Features
- **Custom Output Formats and Widths:** Tailor your image outputs to suit your specific needs.
- **Dark and Light Mode:** Adaptable to your preferred interface style.
- **Target Resolution in Output File:** Easily include your target resolution in the file name
- **Customizable Output Code:** Get the code snippets you need, tailored to your preferences.
- **Retro Design:** Inspired in the old macOS style.

![Wibo in Action](https://raw.githubusercontent.com/eneru2/wibo/main/.github/media/use_case.mp4)


Install from package
-------------------
A pre-built package for Windows can be found on the (this bundles ffmpeg with the app)
[Releases](https://github.com/Eneru2/pixel-morph/releases/) page. In the near future I'll be releasing pre-build packages
for Linux and maybe MacOS, aswell as a CLI-tool geared towards fast and bulk scripting.

Install from source
-------------------
You need to clone the repository:
  ```
  git clone https://github.com/eneru2/wibo
  ```
You will need [npm](https://nodejs.org/en/download),
[Go 1.20+](https://go.dev/dl/) and
[Wails](https://wails.io/docs/gettingstarted/installation/)

Then run the following command after cloning:
  ```
  cd wibo
  wails dev
  ```
After that you will have the executable file corresponding to your OS and architecture.

Future of the project
--------------------
This initial release of Wibo is just the beginning. While it's already functional, there are plans for further enhancements. Future releases will introduce more customization options, allowing users to tweak default values and tailor the tool to their specific use cases. Your feedback and contributions will be invaluable in shaping the future of Wibo.