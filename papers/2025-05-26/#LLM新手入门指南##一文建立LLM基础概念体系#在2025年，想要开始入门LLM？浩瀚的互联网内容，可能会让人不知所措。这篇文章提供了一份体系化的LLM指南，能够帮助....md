# #LLM新手入门指南##一文建立LLM基础概念体系#在2025年，想要开始入门LLM？浩瀚的互联网内容，可能会让人不知所措。这篇文章提供了一份体系化的LLM指南，能够帮助...

**URL**: https://weibo.com/6105753431/PtCxRiE2h

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23LLM%E6%96%B0%E6%89%8B%E5%85%A5%E9%97%A8%E6%8C%87%E5%8D%97%23&amp;extparam=%23LLM%E6%96%B0%E6%89%8B%E5%85%A5%E9%97%A8%E6%8C%87%E5%8D%97%23" data-hide=""><span class="surl-text">#LLM新手入门指南#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E6%96%87%E5%BB%BA%E7%AB%8BLLM%E5%9F%BA%E7%A1%80%E6%A6%82%E5%BF%B5%E4%BD%93%E7%B3%BB%23&amp;extparam=%23%E4%B8%80%E6%96%87%E5%BB%BA%E7%AB%8BLLM%E5%9F%BA%E7%A1%80%E6%A6%82%E5%BF%B5%E4%BD%93%E7%B3%BB%23" data-hide=""><span class="surl-text">#一文建立LLM基础概念体系#</span></a><br><br>在2025年，想要开始入门LLM？<br><br>浩瀚的互联网内容，可能会让人不知所措。这篇文章提供了一份体系化的LLM指南，能够帮助你更快了解LLM【图1】<br><br>一、智能体<br><br>关于智能体的定义有很多，核心在于：自主运行以执行任务、决策后续行动方案、调用工具达成目标。<br><br>一个智能体的关键组成部分包括：模型、指令/编排和工具集。接下来，我们逐一探讨<br><br>二、模型<br><br>这里主要指驱动智能体的语言模型：<br><br>- 大语言模型：拥有数千亿甚至数万亿参数，需在海量数据上训练，消耗巨大计算资源。<br><br>- 小语言模型：专注于特定任务，在较小数据集上训练，平衡了性能与资源效率。<br><br>三、指令/编排<br><br>指令是智能体行为的指导原则与防护机制。为应对复杂场景，智能体常需编排，包含三个层级：<br><br>1. 指令：设定模型目标、角色、行为规则及其他优化信息。<br><br>2. 记忆：通过提示词（prompt）传递给模型的上下文内容。<br><br>- 短期记忆：模型在交互过程中能够获取的即时上下文信息<br><br>- 长期记忆：突破上下文窗口限制，通常会通过外部存储机制调用的重要信息。<br><br>3. 基于模型的推理/规划：推理技术，结构化地获取信息、执行内部推理并做出决策。<br><br>- 思维链（CoT）：要求模型逐步生成解释或推理过程，避免跳过中间步骤导致推理失误。<br><br>- ReAct：使智能体能对用户查询进行推理并采取行动，持续循环直至任务完成。<br><br>四、工具<br><br>借助工具，模型可以与外部系统交互，并获取训练数据以外的知识。<br><br>1. 函数和函数调用：函数是可重复使用的代码模块。模型根据预定义函数的规格说明，判断何时使用及所需参数，但不自行执行函数。你需另外编写代码来实际执行这些函数。<br><br>2. 外部数据：无需重新训练或微调模型，即可提供额外数据。最常见的途径之一是通过检索增强生成（RAG）：【图3】<br>  <br>  - 检索（Retrieval）：LLM接收查询时，RAG系统从外部知识源搜索并获取相关信息。<br><br>  - 增强（Augmented）：将检索到的相关信息整合到原始提示中，形成增强版输入。<br><br>  - 生成（Generation）：LLM基于原始提示和新增的检索上下文，生成最终回答。<br><br>五、增强模型性能<br><br>1. 上下文学习：直接在提示词中提供示例，无需改变模型底层权重，即可“教会”模型执行任务。<br><br>2. 基于检索的上下文学习：模型检索外部上下文（如文档），并将其添加到提示中，以增强响应。<br><br>3. 基于微调的学习（Fine-tuning based Learning）：在特定数据集上进一步训练模型，使其“内化”新的行为或知识。这包括全参数微调、LoRA和RLHF等方法。【图4】<br><br>六、LangChain 的作用是什么？<br><br>LangChain是一个专为简化基于大语言模型应用开发而设计的框架，包含：<br><br>- LangChain ：作为LLM的基础连接层，它支持无缝切换不同服务商或组合功能模块，无需修改底层代码，同时还能简化整体代码结构。<br><br>- LangGraph: 专注于智能体工作流的构建、部署与管理<br><br>- LangSmith: 提供LLM应用的调试、测试及监控能力<br><br>原文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Ftowardsdatascience.com%2Fnew-to-llms-start-here%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1szyb8j9yj37ex3xnb2a.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1szy9sc9wj30cz0bq406.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1szybrfmpj30sg0g00va.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1szydmhj9j30jp04r75e.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

这篇LLM入门指南系统性地介绍了大语言模型的核心概念：1) 智能体由模型、指令编排和工具集构成；2) 模型分为参数量巨大的LLM和轻量级SLM；3) 指令编排包含目标设定、记忆管理（短期/长期）和推理技术（思维链/ReAct）；4) 工具扩展包括函数调用和RAG检索增强；5) 性能优化方法有上下文学习、微调技术（全参数/LoRA/RLHF）；6) LangChain框架简化了LLM应用开发。全文为初学者构建了完整的LLM知识体系框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T23:03:09Z
- **目录日期**: 2025-05-26
