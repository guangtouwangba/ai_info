# TokenHSI: Unified Synthesis of Physical Human-Scene Interactions through Task Tokenization

**URL**: http://arxiv.org/abs/2503.19901v1

## 原始摘要

Synthesizing diverse and physically plausible Human-Scene Interactions (HSI)
is pivotal for both computer animation and embodied AI. Despite encouraging
progress, current methods mainly focus on developing separate controllers, each
specialized for a specific interaction task. This significantly hinders the
ability to tackle a wide variety of challenging HSI tasks that require the
integration of multiple skills, e.g., sitting down while carrying an object. To
address this issue, we present TokenHSI, a single, unified transformer-based
policy capable of multi-skill unification and flexible adaptation. The key
insight is to model the humanoid proprioception as a separate shared token and
combine it with distinct task tokens via a masking mechanism. Such a unified
policy enables effective knowledge sharing across skills, thereby facilitating
the multi-task training. Moreover, our policy architecture supports variable
length inputs, enabling flexible adaptation of learned skills to new scenarios.
By training additional task tokenizers, we can not only modify the geometries
of interaction targets but also coordinate multiple skills to address complex
tasks. The experiments demonstrate that our approach can significantly improve
versatility, adaptability, and extensibility in various HSI tasks. Website:
https://liangpan99.github.io/TokenHSI/


## AI 摘要

TokenHSI是一个基于Transformer的统一策略模型，旨在解决多样化人-场景交互（HSI）任务。传统方法依赖多个独立控制器，难以整合多技能（如边坐下边持物）。TokenHSI通过将人体本体感觉建模为共享令牌，并结合任务令牌的掩码机制，实现跨技能知识共享和多任务训练。其架构支持可变长度输入，使技能灵活适应新场景。通过额外任务令牌化器，可调整交互目标的几何形状并协调多技能处理复杂任务。实验表明，该方法显著提升了HSI任务的通用性、适应性和扩展性。详情见项目网站。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T08:00:59Z
- **目录日期**: 2025-03-26
