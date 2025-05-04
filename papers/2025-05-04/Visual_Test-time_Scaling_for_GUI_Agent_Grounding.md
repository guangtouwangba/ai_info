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

我们提出了RegionFocus，一种用于视觉语言模型代理的视觉测试时缩放方法。该方法通过动态放大相关区域来减少网页GUI图像的视觉复杂性，提高动作选择的准确性。我们采用图像地图机制可视化关键地标，提供透明的操作记录，帮助代理有效选择候选动作。实验显示，该方法在Screenspot-pro和WebVoyager基准测试上分别提升了28%和24%的性能，使Qwen2.5-VL-72B模型在ScreenSpot-Pro上达到61.6%的最新准确率。代码将开源发布。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-04T18:01:08Z
- **目录日期**: 2025-05-04
