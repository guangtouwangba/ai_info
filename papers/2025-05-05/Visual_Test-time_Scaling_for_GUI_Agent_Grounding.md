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

本文提出了RegionFocus，一种视觉测试时缩放方法，用于提升视觉语言模型代理处理网页的能力。该方法通过动态聚焦相关区域减少背景干扰，并采用图像映射机制可视化关键地标，提高动作选择的透明度和准确性。实验表明，在UI-TARS和Qwen2.5-VL等先进模型基础上，该方法在ScreenSpot-pro和WebVoyager基准测试中分别提升28%和24%的性能。特别是将RegionFocus应用于Qwen2.5-VL-72B模型后，在ScreenSpot-Pro基准测试中达到了61.6%的最新最优性能。代码将开源在GitHub上。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T01:28:48Z
- **目录日期**: 2025-05-05
