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

我们提出了RegionFocus，一种用于视觉语言模型代理的视觉测试时缩放方法。该方法通过动态放大网页相关区域来减少背景干扰，提高定位准确性。我们还引入了图像地图机制，可视化关键地标以增强动作选择的透明度。实验表明，即使在简单区域选择策略下，该方法在UI-TARS和Qwen2.5-VL等先进模型基础上，将ScreenSpot-pro和WebVoyager基准测试性能分别提升28%和24%。使用Qwen2.5-VL-72B模型时，ScreenSpot-Pro基准测试达到61.6%的最新定位性能。代码将开源在GitHub上。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-02T22:01:10Z
- **目录日期**: 2025-05-02
