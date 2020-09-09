package dashboard

var (
	//TODO: we should setup static file generation pipeline and inject this variable in building
	FrontendCSS = "https://simplemvc-cdn.oss-cn-shanghai.aliyuncs.com/oam/dist/umi.126737de.css"
	FrontendJS  = "https://simplemvc-cdn.oss-cn-shanghai.aliyuncs.com/oam/dist/umi.06bf354b.js"
)

var IndexHtml = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"
    />
    <title>Micro App Engine</title>
    <!-- <link rel="icon" href="https://simplemvc-cdn.oss-cn-shanghai.aliyuncs.com/oam/dist/favicon.png" type="image/x-icon" /> -->
    <link rel="stylesheet" href="` + FrontendCSS + `" />
    <script>
      window.routerBase = "/";
    </script>
    <script>
      //! umi version: 3.2.18
    </script>
  </head>
  <body>
    <noscript>Out-of-the-box mid-stage front/design solution!</noscript>
    <div id="root">
      <style>
        html,
        body,
        #root {
          height: 100%;
          margin: 0;
          padding: 0;
        }
        #root {
          background-image: url("https://simplemvc-cdn.oss-cn-shanghai.aliyuncs.com/oam/dist/home_bg.png");
          background-repeat: no-repeat;
          background-size: 100% auto;
        }
        .page-loading-warp {
          padding: 98px;
          display: flex;
          justify-content: center;
          align-items: center;
        }
        .ant-spin {
          -webkit-box-sizing: border-box;
          box-sizing: border-box;
          margin: 0;
          padding: 0;
          color: rgba(0, 0, 0, 0.65);
          font-size: 14px;
          font-variant: tabular-nums;
          line-height: 1.5;
          list-style: none;
          -webkit-font-feature-settings: "tnum";
          font-feature-settings: "tnum";
          position: absolute;
          display: none;
          color: #1890ff;
          text-align: center;
          vertical-align: middle;
          opacity: 0;
          -webkit-transition: -webkit-transform 0.3s
            cubic-bezier(0.78, 0.14, 0.15, 0.86);
          transition: -webkit-transform 0.3s
            cubic-bezier(0.78, 0.14, 0.15, 0.86);
          transition: transform 0.3s cubic-bezier(0.78, 0.14, 0.15, 0.86);
          transition: transform 0.3s cubic-bezier(0.78, 0.14, 0.15, 0.86),
            -webkit-transform 0.3s cubic-bezier(0.78, 0.14, 0.15, 0.86);
        }

        .ant-spin-spinning {
          position: static;
          display: inline-block;
          opacity: 1;
        }

        .ant-spin-dot {
          position: relative;
          display: inline-block;
          font-size: 20px;
          width: 20px;
          height: 20px;
        }

        .ant-spin-dot-item {
          position: absolute;
          display: block;
          width: 9px;
          height: 9px;
          background-color: #1890ff;
          border-radius: 100%;
          -webkit-transform: scale(0.75);
          -ms-transform: scale(0.75);
          transform: scale(0.75);
          -webkit-transform-origin: 50% 50%;
          -ms-transform-origin: 50% 50%;
          transform-origin: 50% 50%;
          opacity: 0.3;
          -webkit-animation: antSpinMove 1s infinite linear alternate;
          animation: antSpinMove 1s infinite linear alternate;
        }

        .ant-spin-dot-item:nth-child(1) {
          top: 0;
          left: 0;
        }

        .ant-spin-dot-item:nth-child(2) {
          top: 0;
          right: 0;
          -webkit-animation-delay: 0.4s;
          animation-delay: 0.4s;
        }

        .ant-spin-dot-item:nth-child(3) {
          right: 0;
          bottom: 0;
          -webkit-animation-delay: 0.8s;
          animation-delay: 0.8s;
        }

        .ant-spin-dot-item:nth-child(4) {
          bottom: 0;
          left: 0;
          -webkit-animation-delay: 1.2s;
          animation-delay: 1.2s;
        }

        .ant-spin-dot-spin {
          -webkit-transform: rotate(45deg);
          -ms-transform: rotate(45deg);
          transform: rotate(45deg);
          -webkit-animation: antRotate 1.2s infinite linear;
          animation: antRotate 1.2s infinite linear;
        }

        .ant-spin-lg .ant-spin-dot {
          font-size: 32px;
          width: 32px;
          height: 32px;
        }

        .ant-spin-lg .ant-spin-dot i {
          width: 14px;
          height: 14px;
        }

        @media all and (-ms-high-contrast: none), (-ms-high-contrast: active) {
          .ant-spin-blur {
            background: #fff;
            opacity: 0.5;
          }
        }

        @-webkit-keyframes antSpinMove {
          to {
            opacity: 1;
          }
        }

        @keyframes antSpinMove {
          to {
            opacity: 1;
          }
        }

        @-webkit-keyframes antRotate {
          to {
            -webkit-transform: rotate(405deg);
            transform: rotate(405deg);
          }
        }

        @keyframes antRotate {
          to {
            -webkit-transform: rotate(405deg);
            transform: rotate(405deg);
          }
        }
      </style>
      <div
        style="
          display: flex;
          justify-content: center;
          align-items: center;
          flex-direction: column;
          min-height: 420px;
          height: 100%;
        "
      >
        <div class="page-loading-warp">
          <div class="ant-spin ant-spin-lg ant-spin-spinning">
            <span class="ant-spin-dot ant-spin-dot-spin"
              ><i class="ant-spin-dot-item"></i><i class="ant-spin-dot-item"></i
              ><i class="ant-spin-dot-item"></i><i class="ant-spin-dot-item"></i
            ></span>
          </div>
        </div>
      </div>
    </div>

    <script src="` + FrontendJS + `"></script>
  </body>
</html>`
