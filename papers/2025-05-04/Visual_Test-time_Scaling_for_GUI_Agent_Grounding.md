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

RegionFocus是一种视觉测试时缩放方法，用于提升视觉语言模型代理处理网页的能力。针对GUI图像的复杂性和大量界面元素导致的动作选择困难，该方法动态聚焦相关区域，减少背景干扰，提高定位准确性。通过图像地图机制可视化关键地标，提供透明的动作记录，帮助代理有效选择候选动作。实验表明，该方法在Screenspot-pro和WebVoyager基准测试中分别提升28%和24%的性能，结合Qwen2.5-VL-72B模型时，ScreenSpot-Pro基准测试达到61.6%的最新定位性能。代码将公开于https://github.com/tiangeluo/RegionFocus。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-04T03:21:23Z
- **目录日期**: 2025-05-04
