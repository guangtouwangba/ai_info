# Modeling crystal defects using defect-informed neural networks

**URL**: http://arxiv.org/abs/2503.15391v1

## 原始摘要

Machine learning has revolutionized the study of crystalline materials for
enabling rapid predictions and discovery. However, most AI-for-Materials
research to date has focused on ideal crystals, whereas real-world materials
inevitably contain defects that play a critical role in modern functional
technologies. The defects and dopants break geometric symmetry and increase
interaction complexity, posing particular challenges for traditional ML models.
Addressing these challenges requires models that are able to capture sparse
defect-driven effects in crystals while maintaining adaptability and precision.
Here, we introduce Defect-Informed Equivariant Graph Neural Network (DefiNet),
a model specifically designed to accurately capture defect-related interactions
and geometric configurations in point-defect structures. Trained on 14,866
defect structures, DefiNet achieves highly accurate structural predictions in a
single step, avoiding the time-consuming iterative processes in modern ML
relaxation models and possible error accumulation from iteration. We further
validates DefiNet's accuracy by using density functional theory (DFT)
relaxation on DefiNet-predicted structures. For most defect structures,
regardless of defect complexity or system size, only 3 ionic steps are required
to reach the DFT-level ground state. Finally, comparisons with scanning
transmission electron microscopy (STEM) images confirm DefiNet's scalability
and extrapolation beyond point defects, positioning it as a groundbreaking tool
for defect-focused materials research.


## AI 摘要

机器学习在晶体材料研究中实现了快速预测和发现，但现有研究多集中于理想晶体，忽略了实际材料中的缺陷。这些缺陷和掺杂破坏了几何对称性，增加了相互作用的复杂性，对传统机器学习模型构成挑战。为此，研究者提出了缺陷信息等变图神经网络（DefiNet），专门用于捕捉点缺陷结构中的缺陷相关相互作用和几何配置。DefiNet在14,866个缺陷结构上训练，能一步实现高精度结构预测，避免了迭代过程中的误差积累。通过密度泛函理论验证，DefiNet预测的结构仅需3次离子步骤即可达到基态，展示了其在缺陷材料研究中的突破性潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T13:02:10+08:00
- **目录日期**: 2025-03-20
