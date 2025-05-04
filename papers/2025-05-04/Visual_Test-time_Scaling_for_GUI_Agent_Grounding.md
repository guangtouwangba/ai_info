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

RegionFocus是一种视觉测试时缩放方法，用于提升视觉语言模型代理在网页理解任务中的表现。该方法通过动态放大相关区域来减少背景干扰，提高动作选择的准确性，并采用图像-地图机制可视化关键地标，增强动作记录的透明度。实验表明，在UI-TARS和Qwen2.5-VL等先进模型基础上，RegionFocus在ScreenSpot-Pro和WebVoyager基准测试中分别提升了28%和24%的性能。结合Qwen2.5-VL-72B模型后，ScreenSpot-Pro基准的准确率达到61.6%，刷新了当前最优水平。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-04T08:01:06Z
- **目录日期**: 2025-05-04
