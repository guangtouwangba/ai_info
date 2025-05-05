# Evaluating and Aligning CodeLLMs on Human Preference网页链接这篇论文研究了代码生成大型语言模型（codeLLMs）的性能评估和与人类偏好对齐的问题。现有的大部...

**URL**: https://weibo.com/1870858943/P58h85y3u

## 原始摘要

Evaluating and Aligning CodeLLMs on Human Preference<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F67565a4bae8580e7ff8e0fbd%2F%3Ff%3Dwb" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>这篇论文研究了代码生成大型语言模型（codeLLMs）的性能评估和与人类偏好对齐的问题。现有的大部分代码相关基准测试主要关注生成正确代码片段的能力，但忽略了与人类偏好的匹配。为此，论文提出了一种严格的人类策划的基准CodeArena，模拟现实世界编程任务的复杂性和多样性，并从用户查询中精心选取了397个高质量样本，涵盖了40个类别和44种编程语言。此外，论文还构建了一个多样化的合成指令语料库SynCode-Instruct（近20B个标记），通过扩展网站上的指令来验证大规模合成指令微调的有效性。实验表明，基于执行的基准测试与CodeArena之间存在性能差异，且在40多个大型语言模型上的系统实验显示，开源的顶级代码LLM（如Qwen2.5-Coder）与专有LLM（如OpenAI o1）之间存在显著的性能差距，突显了与人类偏好对齐的重要性。<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#语言模型#</span></a><a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><a href="https://m.weibo.cn/p/index?extparam=%E7%A7%91%E7%A0%94&amp;containerid=100808a62e87d21630c0abf068bf92641e88be" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">科研</span></a><a href="https://m.weibo.cn/p/index?extparam=%E7%A0%94%E7%A9%B6%E7%94%9F&amp;containerid=100808951b26a4a4b4ec29f6ca7f280fccb863" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">研究生</span></a><img style="" src="https://tvax4.sinaimg.cn/large/6f830abfly1hwmt504ez9j22bt17pnpd.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

这篇论文针对代码生成大模型(codeLLMs)提出两个关键贡献：1) 构建CodeArena基准测试，包含397个真实编程任务样本(覆盖40类/44种语言)，首次系统评估模型输出与人类偏好的匹配度；2) 创建20B标记量的SynCode-Instruct合成指令集验证微调效果。实验发现：现有基于执行的评测与人类偏好存在显著差异，开源顶级模型(如Qwen2.5-Coder)与商业模型(如OpenAI o1)在人性化输出方面差距明显。研究强调了代码生成模型除功能性外，还需重视与开发者使用习惯的对齐。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T21:05:04Z
- **目录日期**: 2025-05-05
