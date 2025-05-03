# Visual Test-time Scaling for GUI Agent Grounding

**URL**: http://arxiv.org/abs/2505.00684v1

## 原始摘要

We introduce RegionFocus, a visual test-time scaling approach for Vision
Language Model Agents. Understanding webpages is challenging due to the visual
complexity of GUI images and the large number of interface elements, making
accurate action selection difficult. Our approach dynamically zooms in on
relevant regions, reducing background clutter and improving grounding accuracy.
To support this process, we propose an image-as-map mechanism that visualizes
key landmarks at each step, providing a transparent action record and enables
the agent to effectively choose among action candidates. Even with a simple
region selection strategy, we observe significant performance gains of 28+\% on
Screenspot-pro and 24+\% on WebVoyager benchmarks on top of two
state-of-the-art open vision language model agents, UI-TARS and Qwen2.5-VL,
highlighting the effectiveness of visual test-time scaling in interactive
settings. We achieve a new state-of-the-art grounding performance of 61.6\% on
the ScreenSpot-Pro benchmark by applying RegionFocus to a Qwen2.5-VL-72B model.
Our code will be released publicly at https://github.com/tiangeluo/RegionFocus.


## AI 摘要

研究人员提出了一种名为RegionFocus的视觉测试时缩放方法，用于提升视觉语言模型代理在网页理解任务中的性能。该方法通过动态放大相关区域来减少背景干扰，并引入图像地图机制可视化关键地标，从而提高动作选择的准确性。实验表明，该方法在UI-TARS和Qwen2.5-VL两种先进模型上实现了显著提升：在ScreenSpot-Pro和WebVoyager基准测试中分别获得28%和24%的性能增益。特别是将RegionFocus应用于Qwen2.5-VL-72B模型时，在ScreenSpot-Pro基准测试中达到了61.6%的最新最优性能。相关代码将开源发布。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-03T03:15:37Z
- **目录日期**: 2025-05-03
