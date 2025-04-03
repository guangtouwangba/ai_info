# Gen-C: Populating Virtual Worlds with Generative Crowds

**URL**: http://arxiv.org/abs/2504.01924v1

## 原始摘要

Over the past two decades, researchers have made significant advancements in
simulating human crowds, yet these efforts largely focus on low-level tasks
like collision avoidance and a narrow range of behaviors such as path following
and flocking. However, creating compelling crowd scenes demands more than just
functional movement-it requires capturing high-level interactions between
agents, their environment, and each other over time. To address this issue, we
introduce Gen-C, a generative model to automate the task of authoring
high-level crowd behaviors. Gen-C bypasses the labor-intensive and challenging
task of collecting and annotating real crowd video data by leveraging a large
language model (LLM) to generate a limited set of crowd scenarios, which are
subsequently expanded and generalized through simulations to construct
time-expanded graphs that model the actions and interactions of virtual agents.
Our method employs two Variational Graph Auto-Encoders guided by a condition
prior network: one dedicated to learning a latent space for graph structures
(agent interactions) and the other for node features (agent actions and
navigation). This setup enables the flexible generation of dynamic crowd
interactions. The trained model can be conditioned on natural language,
empowering users to synthesize novel crowd behaviors from text descriptions. We
demonstrate the effectiveness of our approach in two scenarios, a University
Campus and a Train Station, showcasing its potential for populating diverse
virtual environments with agents exhibiting varied and dynamic behaviors that
reflect complex interactions and high-level decision-making patterns.


## AI 摘要

过去20年，人群模拟研究主要聚焦低层任务（如避障）和简单行为（如路径跟随）。为生成更真实的人群场景，本研究提出Gen-C模型，通过大语言模型生成初始场景，再经模拟扩展构建时空交互图。模型采用双变分图自编码器，分别学习图结构（交互）和节点特征（行为）的潜在空间，支持基于自然语言的条件生成。实验在大学校园和火车站场景中验证了该模型能生成反映高层决策的多样化动态交互。该方法避免了真实视频数据标注的繁琐，为虚拟环境提供了灵活的人群行为合成方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T10:01:40Z
- **目录日期**: 2025-04-03
