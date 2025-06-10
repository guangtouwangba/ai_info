# $τ^2$-Bench: Evaluating Conversational Agents in a Dual-Control Environment

**URL**: http://arxiv.org/abs/2506.07982v1

## 原始摘要

Existing benchmarks for conversational AI agents simulate single-control
environments, where only the AI agent can use tools to interact with the world,
while the user remains a passive information provider. This differs from
real-world scenarios like technical support, where users need to actively
participate in modifying the state of the (shared) world. In order to address
this gap, we introduce $\tau^2$-bench, with four key contributions:
  1) A novel Telecom dual-control domain modeled as a Dec-POMDP, where both
agent and user make use of tools to act in a shared, dynamic environment that
tests both agent coordination and communication,
  2) A compositional task generator that programmatically creates diverse,
verifiable tasks from atomic components, ensuring domain coverage and
controlled complexity,
  3) A reliable user simulator tightly coupled with the environment, whose
behavior is constrained by tools and observable states, improving simulation
fidelity,
  4) Fine-grained analysis of agent performance through multiple ablations
including separating errors arising from reasoning vs
communication/coordination.
  In particular, our experiments show significant performance drops when agents
shift from no-user to dual-control, highlighting the challenges of guiding
users. Overall, $\tau^2$-bench provides a controlled testbed for agents that
must both reason effectively and guide user actions.


## AI 摘要

现有对话AI的基准测试通常模拟单控制环境，只有AI能使用工具与世界互动，而用户仅被动提供信息，这与现实场景（如技术支持）不符。为此，研究者提出了τ²-bench，包含四大创新：1) 将电信领域建模为Dec-POMDP双控制环境，用户和AI共同使用工具；2) 通过组合原子组件生成多样化任务；3) 开发与工具和状态绑定的高保真用户模拟器；4) 细粒度分析AI在推理与协作中的表现。实验显示，双控制下AI性能显著下降，突显指导用户的挑战。该基准为需协同用户的AI提供了可控测试平台。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T03:23:32Z
- **目录日期**: 2025-06-10
